module example.com/investment/encoding

go 1.23.4

replace example.com/investment => ..

replace example.com/investment/encoding/csv => ./csv

require (
	example.com/investment v0.0.0-00010101000000-000000000000
	example.com/investment/encoding/csv v0.0.0-00010101000000-000000000000
)
