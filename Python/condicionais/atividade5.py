# Entrar via teclado com três valores distintos.
# Exibir o maior deles.

n1 = int(input("Digite o segundo valor: "))
n2 = int(input("Digite o segundo valor: "))
n3 = int(input("Digite o segundo valor: "))

if n1 > n2 and n1 > n3:
    maior_valor = n1
    print("O maior valor é o", maior_valor)
elif n2 > n1 and n2 > n3:
    maior_valor = n2
    print("O maior valor é o", maior_valor)
elif n3 > n2 and n3 > n1:
    maior_valor = n3
    print("O maior valor é o", maior_valor)
else:
    print("Os numeros são identicos:", n1, n2, n3)