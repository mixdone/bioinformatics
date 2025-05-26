package main

import (
	"fmt"
	"os"

	"github.com/scipipe/scipipe"
)

func main() {
	if len(os.Args) != 3 {
		panic("Usage: go run main.go reads.fastq reference.fna")
	}
	fastqPath := os.Args[1]
	refPath := os.Args[2]

	fmt.Println("FASTQ:", fastqPath)
	fmt.Println("Reference:", refPath)

	wf := scipipe.NewWorkflow("minimap2-nanopore-pipeline", 4)

	readSrc := wf.NewProc("read_src", "")
	readSrc.SetOut("fastq", fastqPath)

	refSrc := wf.NewProc("ref_src", "")
	refSrc.SetOut("ref", refPath)

	fastqc := wf.NewProc("fastqc", "fastqc {i:fastq} -o fastqc_out > {o:log} 2>&1")
	fastqc.SetOut("html", "fastqc_out/{i:fastq|%.fastq|basename}_fastqc.html")
	fastqc.SetOut("log", "logs/fastqc.log")
	fastqc.In("fastq").From(readSrc.Out("fastq"))

	index := wf.NewProc("index", "minimap2 -d {o:mmi} {i:ref} > {o:log} 2>&1 && touch {o:done}")
	index.SetOut("log", "logs/minimap2_index.log")
	index.SetOut("done", "ref_index.done")
	index.SetOut("mmi", "ref.mmi")
	index.In("ref").From(refSrc.Out("ref"))

	// -------- Alignment --------
	align := wf.NewProc("align", "minimap2 -ax map-ont {i:index} {i:reads} > {o:sam} 2> {o:log}")
	align.SetOut("sam", "aligned.sam")
	align.SetOut("log", "logs/align.log")
	align.In("reads").From(readSrc.Out("fastq"))
	align.In("index").From(index.Out("mmi"))

	sam2bam := wf.NewProc("sam2bam", "samtools view -b {i:sam} > {o:bam} 2> {o:log}")
	sam2bam.SetOut("bam", "aligned.bam")
	sam2bam.SetOut("log", "logs/sam2bam.log")
	sam2bam.In("sam").From(align.Out("sam"))

	flagstat := wf.NewProc("flagstat", "samtools flagstat {i:bam} > {o:report} 2> {o:log}")
	flagstat.SetOut("report", "alignment_report.txt")
	flagstat.SetOut("log", "logs/flagstat.log")
	flagstat.In("bam").From(sam2bam.Out("bam"))

	evaluate := wf.NewProc("evaluate", `bash -c '
	mapped_percent=$(grep -oE "[0-9]+\.[0-9]+%" -m1 {i:report} | tr -d "%")
	echo "Mapped percent: $mapped_percent%"
	if (( $(echo "$mapped_percent > 90" | bc -l) )); then
		echo "OK"
	else
		echo "Not OK"
	fi ' > {o:log} 2>&1`)
	evaluate.SetOut("log", "logs/evaluate.log")
	evaluate.In("report").From(flagstat.Out("report"))

	wf.PlotGraphPDF("graph/result.pdf")

	wf.Run()
}
