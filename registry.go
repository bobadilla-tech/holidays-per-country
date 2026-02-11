package holidays

var registry = map[string]provider{}

// TODO(crydafan): Cache holidays for each country and year to avoid redundant calculations
