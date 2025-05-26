#!/bin/bash


mapped_percent=$(grep "mapped (" | head -1 | awk '{print $5}' | tr -d '()%:')

if [[ -z "$mapped_percent" ]]; then
    echo "Нет нужных данных" >&2
    exit 1
fi

echo "Mapped reads: $mapped_percent%"

pivot=90

mapped_int=${mapped_percent%.*}

if [ "$mapped_int" -ge "$pivot" ]; then
    echo "OK"
else
    echo "not OK"
fi
