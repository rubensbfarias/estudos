# Criar uma rotina ou função de entrada de dados que aceite somente
# um valor positico (>0). Ou seja, enquanto o usuario não entrar com um valor
# positivo, deve ser requisitado a digitar novamente.

n = int(input("Digite um número positivo: "))

while True:
    if n < 0:
         n = int(input("Ops! Apenas positivos, tente novamente: "))
    else:
         print("Programa finalizado!")
         break       
