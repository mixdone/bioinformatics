#  ALBINISM, OCULOCUTANEOUS, TYPE II; OCA2 

Tyrosinase-положительный глазокожный альбинизм (OCA, тип II; OCA2) представляет собой аутосомно-рецессивное расстройство, при котором биосинтез пигмента меланина уменьшается в коже, волосах и глазах. Люди с OCA типа II имеют характерные визуальные аномалии, связанные с альбинизмом, включая снижение остроты и нистагмуса, которые обычно менее тяжелые, чем в OCA тип I
OCA Type II имеет сильно варьируемый фенотип. Волосы затронутых людей могут стать темными с возрастом, и можно увидеть пигментированные невизийные или веснушки.

## Ассоциированные гены
-  OCA2
-  MC1R

## Оценка качества выравниваний

- Program: matcher      
  ```
  #=======================================
  #
  # Aligned_sequences: 2
  # 1: 124133839-124137483
  # 2: 89918862-89920972
  # Matrix: EDNAFULL
  # Gap_penalty: 16
  # Extend_penalty: 4
  #
  # Length: 1667
  # Identity:    1194/1667 (71.6%)
  # Similarity:  1194/1667 (71.6%)
  # Gaps:          94/1667 ( 5.6%)
  # Score: 3706
  # 
  #
  #=======================================
  ```
- Program: stretcher    
  ```
  #=======================================
  #
  # Aligned_sequences: 2
  # 1: 124133839-124137483
  # 2: 89918862-89920972
  # Matrix: EDNAFULL
  # Gap_penalty: 16
  # Extend_penalty: 4
  #
  # Length: 4327
  # Identity:    2388/4327 (55.2%)
  # Similarity:  2390/4327 (55.2%)
  # Gaps:         747/4327 (17.3%)
  # Score: 1990
  # 
  #
  #=======================================
  ```

matcher показал себя лучше.

OCA2 оказалась слишком длинной последовательностью поэтому выравнивание для нее не построилось. Чуть позже внесу исправления, добавив какой-нибудь другой ген.
