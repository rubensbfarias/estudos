# Entrada de dados -> idade

idade = int(input("Digite sua idade: "))

if idade > 18 and idade < 65:
   print("Você pode dirigir!")
elif idade < 18:
   print("Você não pode dirigir!")
elif idade >= 65:
   print("Você não pode dirigir!")

if idade == 18:
   print ("Você pode se alistar!")