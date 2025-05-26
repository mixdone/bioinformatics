#!/bin/bash


fastqc SRR33692910_1.fastq SRR33692910_2.fastq

minimap2 -d ecoli_reference.mmi GCA_000005845.2_ASM584v2_genomic.fna
minimap2 -ax sr ecoli_reference.mmi reads/SRR33692910_1.fastq reads/SRR33692910_2.fastq > align.sam

samtools view -Sb align.sam > align.bam
samtools sort align.bam -o align.sorted.bam
samtools index aln.sorted.bam
samtools index align.sorted.bam

samtools flagstat aln.sorted.bam > flagstat.txt
samtools flagstat align.sorted.bam > flagstat.txt

cat flagstat.txt 

bash mapped.sh < flagstat.txt