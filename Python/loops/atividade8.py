# Criar um programa para entrar com dois valores inteiros, onde o segundo valor deverá ser maior que o primeiro valor. Caso contrário solicitar
# novamente apenas o segundo valor. 

n1 = int(input("Digite o primeiro valor: "))
n2 = int(input("Digite o segundo valor: "))

while True:
    if n2 < n1:
        n2 = int(input("Digite o segundo valor novamente: "))
    else:
        print("O segundo valor é maior que o primeiro", n2)
        break
