cor = input("Digite a cor do seu carro:\n")

if cor != "preto" or cor != "vermelho" and cor == "azul" or cor == "branco":
    print("Carro padrão")
elif cor != "azul" or cor != "branco":
    print("Carro fora do padrão")
else:
    print("Cor não aceita")