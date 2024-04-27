create table states(
    id          serial primary key,
    name        varchar(50),
    capital     varchar(50),
    region      varchar(50),
    area        varchar(50),
    population  varchar(50)
);

INSERT INTO states(name, capital, region, area, population) VALUES
('Acre', 'Rio Branco', 'Norte', '164.173,431 km²', '906.876'),
('Alagoas', 'Maceió', 'Nordeste', '27.848,158 km²', '3.127.683'),
('Amapá', 'Macapá', 'Norte', '142.470,162 km²', '733.759'),
('Amazonas', 'Manaus', 'Norte', '1.559.167,878 km²', '3.941.613'),
('Bahia', 'Salvador', 'Nordeste', '564.733,177 km²', '15.203.934'),
('Ceará', 'Fortaleza', 'Nordeste', '148.920,472 km²', '9.187.103'),
('Distrito Federal', 'Brasília', 'Centro-Oeste', '5.761 km²', '3.561.000'),
('Espírito Santo', 'Vitória', 'Sudeste', '46.095,583 km²', '4.064.052'),
('Goiás', 'Goiânia', 'Centro-Oeste', '340.111,783 km²', '7.113.540'),
('Maranhão', 'São Luís', 'Nordeste', '331.983,293 km²', '7.114.598'),
('Mato Grosso', 'Cuiabá', 'Centro-Oeste', '903.357,908 km²', '3.526.220'),
('Mato Grosso do Sul', 'Campo Grande', 'Centro-Oeste', '357.145,532 km²', '2.809.394'),
('Minas Gerais', 'Belo Horizonte', 'Sudeste', '586.528,293 km²', '21.292.666'),
('Pará', 'Belém', 'Norte', '1.247.954,666 km²', '8.690.745'),
('Paraíba', 'João Pessoa', 'Nordeste', '56.469,778 km²', '4.039.277'),
('Paraná', 'Curitiba', 'Sul', '199.307,922 km²', '11.516.840'),
('Pernambuco', 'Recife', 'Nordeste', '98.148 km²', '8.796.448'),
('Piauí', 'Teresina', 'Nordeste', '251.578 km²', '3.118.360'),
('Rio de Janeiro', 'Rio de Janeiro', 'Sudeste', '43.780 km²', '15.989.929'),
('Rio Grande do Norte', 'Natal', 'Nordeste', '52.811 km²', '3.168.027'),
('Rio Grande do Sul', 'Porto Alegre', 'Sul', '281.748,538 km²', '11.422.973'),
('Rondônia', 'Porto Velho', 'Norte', '237.576,167 km²', '1.796.460'),
('Roraima', 'Boa Vista', 'Norte', '224.299,506 km²', '631.181'),
('Santa Catarina', 'Florianópolis', 'Sul', '95.346,181 km²', '7.252.502'),
('São Paulo', 'São Paulo', 'Sudeste', '248.209,426 km²', '46.289.333'),
('Sergipe', 'Aracaju', 'Nordeste', '21.910,354 km²', '2.318.822'),
('Tocantins', 'Palmas', 'Norte', '277.620,914 km²', '1.590.248');




