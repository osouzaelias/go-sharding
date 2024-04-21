#!/bin/bash

# URL para onde as requisições serão enviadas
url="http://localhost:8080/customer"

# Inicializa a string para armazenar CNPJs
used_cnpjs=""

# Função para gerar um CNPJ fictício que não tenha sido usado antes
generate_cnpj() {
  local cnpj
  while : ; do
    cnpj=$((RANDOM % 90000000000000 + 10000000000000))
    if [[ $used_cnpjs != *"$cnpj"* ]]; then
      used_cnpjs+="$cnpj "
      echo $cnpj
      break
    fi
  done
}

# Função para gerar um nome de empresa fictício
generate_company_name() {
  local names=("Empresa" "Corporação" "Grupo" "Indústria" "Comércio")
  local suffixes=("Ltda" "S/A" "e Associados" "Internacional" "do Brasil")
  echo "${names[$RANDOM % ${#names[@]}]} ${suffixes[$RANDOM % ${#suffixes[@]}]}"
}

# Executa a chamada curl 1000 vezes com dados aleatórios
for ((i=0; i<1000; i++))
do
  cnpj=$(generate_cnpj)
  company_name=$(generate_company_name)

  curl -X POST "$url" \
      -H "Content-Type: application/json" \
      -d "{
          \"id\": \"$cnpj\",
          \"name\": \"$company_name\"
      }" > /dev/null 2>&1
done

echo "Data load finished."