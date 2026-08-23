
ALTER TABLE lista_item
ADD CONSTRAINT fk_lista_item
FOREIGN KEY (id_lista) REFERENCES lista(id)
ON DELETE RESTRICT;

ALTER TABLE lista_item
ADD CONSTRAINT fk_item_lista
FOREIGN KEY (id_item) REFERENCES item(id)
ON DELETE RESTRICT;

ALTER TABLE item_categoria
ADD CONSTRAINT fk_item_categoria
FOREIGN KEY (id_item) REFERENCES item(id)
ON DELETE RESTRICT;

ALTER TABLE item_categoria
ADD CONSTRAINT fk_categoria_item
FOREIGN KEY (id_categoria) REFERENCES categoria(id)
ON DELETE RESTRICT;


