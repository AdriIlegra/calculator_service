markdown
Copy code
# Microserviço de Calculadora em Go

Este é um microsserviço simples de calculadora em Go que fornece endpoints para várias operações matemáticas, bem como a capacidade de visualizar o histórico de operações.

## Endpoints Disponíveis

### Soma
/calc/soma/$a/$b

bash
Copy code
Realiza a adição dos valores $a e $b.

### Subtração
/calc/sub/$a/$b

bash
Copy code
Realiza a subtração dos valores $a e $b.

### Multiplicação
/calc/mul/$a/$b

bash
Copy code
Realiza a multiplicação dos valores $a e $b.

### Divisão
/calc/div/$a/$b

bash
Copy code
Realiza a divisão dos valores $a e $b.

### Histórico
/calc/historico

perl
Copy code
Retorna o histórico de todas as operações anteriores.

## Como Executar

Certifique-se de ter o Go instalado em seu ambiente. Para iniciar o servidor, execute o seguinte comando no terminal:

```bash
go run calculator.go
O servidor estará rodando em http://localhost:8080.

Exemplos de Uso
Soma
bash
Copy code
curl http://localhost:8080/calc/soma/5/3
Histórico
bash
Copy code
curl http://localhost:8080/calc/historico