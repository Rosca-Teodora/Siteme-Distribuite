# Lab 1: Go fundamentals

```go mod init [nume modul]``` = creeaza un modul in care se mentioneaza dependentele folosite 

**OBS:** un pachet fara functia ```main``` nu are punct de intrare si, deci, nu este executabil 

### Tipuri de date: 
- uint = unsigned 
- byte = int 8
- rune = alias byte

### Declarare: 
- var = modificabil
- const = NEMODIFICABIL 

Operatorul ```:=``` deduce automat tipul. 


## Error handling
O functie care poate ajunge la o eroare este **obligata** sa returneze si un obiect error. Mereu trebuie verificata aceasta valoarea ca fiind diferita de null. 

```defer``` = inainte de a iesi din functie *apeleaza file.close*