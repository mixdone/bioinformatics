# Задание 3

Oxford Nanopore MinION/GridION/PromethION (ONT)	minimap2 https://github.com/lh3/minimap2 SciPipe https://scipipe.org/ 

### Ссылкf на загруженные прочтения из NCBI SRA

https://www.ncbi.nlm.nih.gov/sra/SRX28236620[accn]?report=Full

### Bash-скрипт

Находятся в папке scripts

### Результат команды samtools ﬂagstat

```
509181 + 0 in total (QC-passed reads + QC-failed reads)
481976 + 0 primary
18926 + 0 secondary
8279 + 0 supplementary
0 + 0 duplicates
0 + 0 primary duplicates
507145 + 0 mapped (99.60% : N/A)
479940 + 0 primary mapped (99.58% : N/A)
0 + 0 paired in sequencing
0 + 0 read1
0 + 0 read2
0 + 0 properly paired (N/A : N/A)
0 + 0 with itself and mate mapped
0 + 0 singletons (N/A : N/A)
0 + 0 with mate mapped to a different chr
0 + 0 with mate mapped to a different chr (mapQ>=5)
```

```
Mapped percent: 99.60%
OK
```

### Визуализация

Scipipe предоставляет возможность визуализации через Graphviz


Все объемные файлы удалены и в репозеторий загружены не были. Остальное можно посмотреть.

## Использование scipipe:

```
go run pipeline.go SRR32962948.fastq GCA_000005845.2_ASM584v2_genomic.fna 
```

## Использование scripts:
```
bash pipeline.sh
```

