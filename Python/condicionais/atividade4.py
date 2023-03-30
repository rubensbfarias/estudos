# Entrara com dois valores quaisquer. Exibir o maior deles
# se existir, caso contrario, enviar mensagem avisando que os numeros são identicos.

n1 = int(input("Digite o primeiro valor:"))
n2 = int(input("Digite o segundo valor: "))

if n1 > n2:
    print("O maior número é o ", n1)
elif n2 > n1:
    print("O maior número é o ", n2)
else:
    print("Os números são identicos: ", n1, n2)