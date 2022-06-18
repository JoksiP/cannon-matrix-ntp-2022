# Predmetni projekat NTP

Predefinisan projekat - Kanonov algoritam za množenje matrica

### Student:
* Jovan Petljanski, SW-31/2018

### Opis algoritma

Kanonov algoritam je distribuiran algoritam za množenje dvodimenzionalnih nizova.
  - Uzmimo dve *N* x *N* matrice *A* i *B* particionisane u *P* blokova
  - *A*(*i*, *j*) i *B*(*i*, *j*) (0 ≤ *i*, *j* ≤ √*P*) veličine (*N* ∕ √*P*)×(*N* ∕ √*P*)
  - Proces *P*(*i*, *j*) inicijalno čuva *A*(*i*, *j*) i *B*(*i*, *j*) i računa blok *C*(*i*, *j*) rezultujuće matrice
  - Inicijalni korak algoritma "pomera" matrice *A* i *B* tako da svaki proces može krenuti sa računom odvojeno
  - Ovo se izvršava tako što se sve vrste matrice *A*(*i*) šiftuju za *i* koraka i kolone matrice *B*(:, *j*) za *j* koraka
  - Lokalni blokovi se množe
  - Ponovo se vrše šiftovi za 1 korak
  - Izvršava se množenje narednog bloka
  - Ponavlja se dok svi blokovi nisu izmnoženi i dodati u rezultujuću matricu *C*

### Opis sekvencijalne implementacije algoritma

Pre iteracionog procesa se izvršava inicijalni šift matrica *A* i *B*. Zatim se u narednih *P* iteracija množe matrice *A* i *B*, i rezultat dodaje u rezultujuću matricu *C*. Na kraju svake iteracije matrice *A* i *B* se šiftuju 1 korak. Konačni rezultat algoritma je matrica *C* = *A* x *B*.

### Načini implementacije algoritama i tehnologije upotrebljene

Algoritam je implentiran u Python i Golang programskim jezicima, paralelno i sekvencijalno. Takođe će se vršiti eksperimenti slabog i jakog skaliranja tako što će se meriti srednje vreme izvršavanja algoritma sekvencijalno i paralelno u oba jezika.

Biblioteka multiprocessing za paralelizaciju je upotrebljena u jeziku Python.

### Vizualizacija rešenja

Rešenje će biti vizualizovano tako što će se prikazati svaka iteracija sekvencijalnog rešenja, i svaki proces i njegov rezultat paralelnog rešenja.
