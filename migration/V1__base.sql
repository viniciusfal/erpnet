
CREATE TABLE lista(
    id UUID PRIMARY KEY NOT NULL
    , nome VARCHAR(100) NOT NULL
    , descricao TEXT
    , data_criacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    , data_atualizacao TIMESTAMP
); 

CREATE TABLE item(
    id UUID PRIMARY KEY NOT NULL
    , nome VARCHAR(100) NOT NULL
    , descricao TEXT
    , quantidade INT
    , url_image TEXT
    , valor_unitario DECIMAL(10,2) NOT NULL
    , data_criacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    , data_atualizacao TIMESTAMP
);

CREATE TABLE lista_item(
    id_lista UUID NOT NULL
    , id_item UUID NOT NULL 
);

CREATE TABLE categoria(
     id BIGINT GENERATED ALWAYS AS IDENTITY
    , nome VARCHAR(70) NOT NULL
); 
   

CREATE TABLE item_categoria(
    id_item UUID NOT NULL
    , id_categoria UUID NOT NULL 
);




