# get beginning of the data set
head(PlantGrowth)

# summarize the data set
str(PlantGrowth)

# column 1
PlantGrowth[, 1]

# row 11 to 20 all columns
PlantGrowth[11:20, ]

# pick variable weight in the data set PlantGrowth
with(PlantGrowth, weight)

# pick file
file.choose()

# and
is_false <- TRUE & FALSE

# or
is_true <- TRUE | FALSE

# use ? to open help
?read.table
