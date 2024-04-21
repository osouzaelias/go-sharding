#!/bin/bash

ENDPOINT_URL="http://localhost:8000"

echo "Configurando o AWS CLI..."
aws configure set default.region us-west-2
aws configure set aws_access_key_id fakeMyKeyId
aws configure set aws_secret_access_key fakeSecretAccessKey

# Lista as tabelas existentes e verifica se cada uma deve ser excluída antes de ser recriada
existing_tables=$(aws dynamodb list-tables --endpoint-url $ENDPOINT_URL --query "TableNames[]" --output text)

echo "Criando as tabelas no DynamoDB Local..."

for i in {1..3}
do
    table_name="node$i"

    # Verifica se a tabela já existe
    if [[ $existing_tables =~ $table_name ]]; then
        echo "A tabela $table_name já existe. Excluindo..."
        aws dynamodb delete-table --endpoint-url $ENDPOINT_URL --table-name $table_name
        echo "Tabela $table_name excluída."

        # Espera a tabela ser excluída antes de tentar recriá-la
        echo "Aguardando a tabela $table_name ser excluída completamente..."
        aws dynamodb wait table-not-exists --endpoint-url $ENDPOINT_URL --table-name $table_name
    fi

    echo "Criando a tabela $table_name..."

    aws dynamodb create-table --endpoint-url $ENDPOINT_URL \
        --table-name $table_name \
        --attribute-definitions \
            AttributeName=id,AttributeType=S \
        --key-schema \
            AttributeName=id,KeyType=HASH \
        --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1

    echo "Tabela $table_name criada."
done

echo "Listando tabelas disponíveis..."
aws dynamodb list-tables --endpoint-url $ENDPOINT_URL