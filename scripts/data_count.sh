#!/bin/bash

ENDPOINT_URL="http://localhost:8000"
TABLES=("node1" "node2" "node3") # Nomes das tabelas

echo "Iniciando a contagem de registros..."

# Itera sobre cada tabela para executar a operação Scan e contar os itens
for table_name in "${TABLES[@]}"
do
    echo "Scaneando a tabela $table_name..."

    # Executa o scan e processa a saída para extrair a contagem de registros
    result=$(aws dynamodb scan --endpoint-url $ENDPOINT_URL \
        --table-name $table_name \
        --select "COUNT")

    # Extrai o valor da contagem usando grep e awk
    count=$(echo "$result" | grep "Count" | awk '{print $2}' | tr -d ',')

    echo "Total de registros na tabela $table_name: $count"
done

echo "Contagem de registros concluída."
