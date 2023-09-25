# La ruta mas corta con Algoritmo A\*

## Tabla de descripcion de conexiones:

| CITY | X     | Y     | CONNECTION(distances)                   |
| ---- | ----- | ----- | --------------------------------------- |
| A    | 0.00  | 0.00  | B (1.41), C (2.83)                      |
| B    | 1.00  | 1.00  | A (1.41), D (7.21), E (2.83)            |
| C    | 2.00  | 2.00  | A (2.83), D (5.83), E (1.41), H (15.56) |
| D    | 5.00  | 7.00  | B (7.21), C (5.83), F (5.10)            |
| E    | 3.00  | 3.00  | B (2.83), C (1.41), G (10.00)           |
| F    | 10.00 | 8.00  | D (5.10), G (5.83)                      |
| G    | 13.00 | 3.00  | E (10.00), F (5.83), H (10.00)          |
| H    | 13.00 | 13.00 | C (15.56), G (10.00)                    |

## Mapa de conexiones:

![alt](./map.png)

## Ejecucion:
Para ejecutar el programa se debe ingresar el comando:
```go run main.go```

