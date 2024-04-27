package models

type State struct {
	Id		   	int	 	`json:"id"`	
    Name       	string 	`json:"name"`
    Capital    	string 	`json:"capital"`
    Region     	string 	`json:"region"`
    Area       	string 	`json:"area"`
    Population 	string 	`json:"population"`
}

var States = []State{
	{
		Id: 1, //ok
		Name:       "Acre",
		Capital:    "Rio Branco",
		Region:     "Norte",
		Area:       "164.173,431 km²",
		Population: "906.876",
	},

	{
		Id: 2, //ok
		Name:       "Alagoas",
		Capital:    "Maceió",
		Region:     "Nordeste",
		Area:       "27.848,158 km²",
		Population: "3.127.683",
	},

	{
		Id: 3, //ok
		Name:       "Amapá",
		Capital:    "Macapá",
		Region:     "Norte",
		Area:       "142.470,162 km²",
		Population: "733.759",
	},

	{
		Id: 4, //ok
		Name:       "Amazonas",
		Capital:    "Manaus",
		Region:     "Norte",
		Area:       "1.559.167,878 km²",
		Population: "3.941.613",
	},

	{
		Id: 5, //ok
		Name:       "Bahia",
		Capital:    "Salvador",
		Region:     "Nordeste",
		Area:       "564.733,177 km²",
		Population: "15.203.934",
	},

	{
		Id: 6, //ok
		Name:       "Ceará",
		Capital:    "Fortaleza",
		Region:     "Nordeste",
		Area:       "148.920,472 km²",
		Population: "9.187.103",
	},

	{
		Id: 7, //ok
		Name:       "Espírito Santo",
		Capital:    "Vitória",
		Region:     "Sudeste",
		Area:       "46.095,583 km²",
		Population: "4.064.052",
	},

	{
		Id: 8, //ok
		Name:       "Goiás",
		Capital:    "Goiânia",
		Region:     "Centro-Oeste",
		Area:       "340.111,783 km²",
		Population: "7.113.540",
	},

	{
		Id: 9, //ok
		Name:       "Maranhão",
		Capital:    "São Luís",
		Region:     "Nordeste",
		Area:       "331.983,293 km²",
		Population: "7.114.598",
	},

	{
		Id: 10, //ok
		Name:       "Mato Grosso",
		Capital:    "Cuiabá",
		Region:     "Centro-Oeste",
		Area:       "903.357,908 km²",
		Population: "3.526.220",
	},

	{
		Id: 11, //ok
		Name:       "Mato Grosso do Sul",
		Capital:    "Campo Grande",
		Region:     "Centro-Oeste",
		Area:       "357.145,532 km²",
		Population: "2.809.394",
	},

	{
		Id: 12, //ok
		Name:       "Minas Gerais",
		Capital:    "Belo Horizonte",
		Region:     "Sudeste",
		Area:       "586.528,293 km²",
		Population: "21.292.666",
	},

	{
		Id: 13, //ok
		Name:       "Pará",
		Capital:    "Belém",
		Region:     "Norte",
		Area:       "1.247.954,666 km²",
		Population: "8.690.745",
	},

	{
		Id: 14, //ok
		Name:       "Paraíba",
		Capital:    "João Pessoa",
		Region:     "Nordeste",
		Area:       "56.469,778 km²",
		Population: "4.039.277",
	},

	{
		Id: 15, //ok
		Name:       "Pernambuco",
		Capital:    "Recife",
		Region:     "Nordeste",
		Area:       "98.148 km²",
		Population: "8.796.448",
	},

	{
		Id: 16, //ok
		Name:       "Piauí",
		Capital:    "Teresina",
		Region:     "Nordeste",
		Area:       "251.578 km²",
		Population: "3.118.360",
	},

	{
		Id: 17, //ok
		Name:       "Rio de Janeiro",
		Capital:    "Rio de Janeiro",
		Region:     "Sudeste",
		Area:       "43.780 km²",
		Population: "15.989.929",
	},

	{
		Id: 18, //ok
		Name:       "Rio Grande do Norte",
		Capital:    "Natal",
		Region:     "Nordeste",
		Area:       "52.811 km²",
		Population: "3.168.027",
	},

	{
		Id: 19, //ok
		Name:       "Rio Grande do Sul",
		Capital:    "Porto Alegre",
		Region:     "Sul",
		Area:       "281.748,538 km²",
		Population: "11.422.973",
	},

	{
		Id: 20, //ok
		Name:       "Rondônia",
		Capital:    "Porto Velho",
		Region:     "Norte",
		Area:       "237.576,167 km²",
		Population: "1.796.460",
	},

	{
		Id: 21, //ok
		Name:       "Roraima",
		Capital:    "Boa Vista",
		Region:     "Norte",
		Area:       "224.299,506 km²",
		Population: "631.181",
	},

	{
		Id: 22, //ok
		Name:       "Santa Catarina",
		Capital:    "Florianópolis",
		Region:     "Sul",
		Area:       "95.346,181 km²",
		Population: "7.252.502",
	},

	{
		Id: 23, //ok
		Name:       "São Paulo",
		Capital:    "São Paulo",
		Region:     "Sudeste",
		Area:       "248.209,426 km²",
		Population: "46.289.333",
	},

	{
		Id: 24, //ok
		Name:       "Sergipe",
		Capital:    "Aracaju",
		Region:     "Nordeste",
		Area:       "21.910,354 km²",
		Population: "2.318.822",
	},

	{
		Id: 25, //ok
		Name:       "Tocantins",
		Capital:    "Palmas",
		Region:     "Norte",
		Area:       "277.620,914 km²",
		Population: "1.590.248",
	},

	{
		Id: 26, //ok
		Name:       "Paraná",
		Capital:    "Curitiba",
		Region:     "Sul",
		Area:       "199.307,92 km²",
		Population: "11.444.380",
	},

	{
		Id: 27, //ok
		Name:       "Distrito Federal",
		Capital:    "Brasília",
		Region:     "Centro-Oeste",
		Area:       "5.761 km²",
		Population: "3.561.000",
	},
}