package delegate

type Add[Lhs, Rhs, Output any] func(Lhs, Rhs) Output

type Sub[Lhs, Rhs, Output any] func(Lhs, Rhs) Output

type Mul[Lhs, Rhs, Output any] func(Lhs, Rhs) Output

type Div[Lhs, Rhs, Output any] func(Lhs, Rhs) Output
