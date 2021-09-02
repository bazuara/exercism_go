package partyrobot

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, "+name+"!"
}

// HappyBirthday wishes happy birthday to the birthday person and stands out his age.
// Output: Happy birthday Frank! You are now 58 years old!
func HappyBirthday(name string, age int) string {
	return "Happy birthday "+name+"! You are now "+strconv.itoa(age)+" years old"
}

// AssignTable assigns a table to each guest.
// Output:
// Welcome to my party, Christiane!
// You have been assigned to table 1B. Your table is on the left, exactly 23.8 meters from here.
// You will be sitting next to Frank.
func AssignTable(name string, table int, neighbour string, direction string, distance float64) string {
	return Welcome(name)+"\n"+
    "You have been assigned to table"+strconv.FormatInt(table, 16)+
    ". Your table is on the "+direction+
    ", exactly "+strconv.FormatFloat(distance, 'f', 1, 64)+"meters from here."
    
}
