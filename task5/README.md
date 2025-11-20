# Задание 5

- Последовательность: MSTLTSVSGFPRIGQNRELKKIIEGYWKGANDLAAVKATAAELRAKHWRLQQAAGIDLIASNDFSYYDQMLDTAILLNVIPQRYQRLAFDDQEDTLFAMA
- Программы-предсказатели: RoseTTAFold2, AlphaFold2
- Программа-выравниватель: TMAlign

## Ноутбуки с предсказаниями

### Copy_of_AlfaFold2.ipynb
Полученное предсказание: alfafold2.pdb
<img width="637" height="481" alt="image" src="https://github.com/user-attachments/assets/d6278207-819a-456f-9790-9a043a80955a" />

### Copy_of_roseTTAFold2.ipynd
Полученное предсказание: rf2.pdb
<img width="637" height="481" alt="image" src="https://github.com/user-attachments/assets/179149c5-edfd-46d2-a311-2c515f9300bc" />

## Вывод программы выравнивания

```
 **************************************************************************
 *                        TM-align (Version 20190822)                     *
 * An algorithm for protein structure alignment and comparison            *
 * Based on statistics:                                                   *
 *       0.0 < TM-score < 0.30, random structural similarity              *
 *       0.5 < TM-score < 1.00, in about the same fold                    *
 * Reference: Y Zhang and J Skolnick, Nucl Acids Res 33, 2302-9 (2005)    *
 * Please email your comments and suggestions to: zhanglab@zhanggroup.org *
 **************************************************************************

Name of Chain_1: A980551                                           
Name of Chain_2: B980551                                           
Length of Chain_1:  100 residues
Length of Chain_2:  100 residues

Aligned length=  100, RMSD=   0.75, Seq_ID=n_identical/n_aligned= 1.000
TM-score= 0.96385 (if normalized by length of Chain_1)
TM-score= 0.96385 (if normalized by length of Chain_2)
(You should use TM-score normalized by length of the reference protein)

(":" denotes aligned residue pairs of d < 5.0 A, "." denotes other aligned residues)
MSTLTSVSGFPRIGQNRELKKIIEGYWKGANDLAAVKATAAELRAKHWRLQQAAGIDLIASNDFSYYDQMLDTAILLNVIPQRYQRLAFDDQEDTLFAMA
::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
MSTLTSVSGFPRIGQNRELKKIIEGYWKGANDLAAVKATAAELRAKHWRLQQAAGIDLIASNDFSYYDQMLDTAILLNVIPQRYQRLAFDDQEDTLFAMA
```
Результат в формате pdb: tm.pdb


## Визуализация выравнивания jsmol
<img width="490" height="464" alt="image" src="https://github.com/user-attachments/assets/53e2450f-7d9f-4504-877b-79a2cd5780a9" />
(Protein-1 in blue and Protein-2 in red)

## Визуализация выравнивания Mol* 3D Viewer

<img width="429" height="335" alt="image" src="https://github.com/user-attachments/assets/ca985a1a-1f23-4795-b08f-f3e7cb7b3f5c" />
(Protein-1 in blue and Protein-2 in red)


## Вывод
По результам выравнивания можно сказать, что полученны предсказания имеют очень сильное сходство, что ясно видно как на визуализациях,так и на выводе выравнивателя.




