# Usuario precisa digitar nome/idade
# As duas ultimas entradas de dados serão nome/idade também porém diferentes.
# 1 - Saida a maior idade
# 2 - Saida a menor idade
# 3 - Saida printar o nome e a idade do primeiro e do segundo

nome1 = input("Digite seu nome: ")
idade1 = int(input("Digite sua idade: "))

nome2 = input("Digite seu nome: ")
idade2 = int(input("Digite sua idade: "))

maior_idade = 0
menor_idade = 0

if idade1 > idade2:
    maior_idade = idade1
    menor_idade = idade2
else:
    maior_idade = idade2
    menor_idade = idade2

print("A maior idade é: ", maior_idade)
print("A menor idade é: ", menor_idade)

print(nome1 + " tem", idade1, "anos")
print(nome2 + " tem", idade2, "anos")   





