package delegate

type BitAnd[Lhs, Rhs, Output any] Function2[Lhs, Rhs, Output]
type BitOr[Lhs, Rhs, Output any] Function2[Lhs, Rhs, Output]
type BitNot[T, Output any] Function1[T, Output]
type BitXor[Lhs, Rhs, Output any] Function2[Lhs, Rhs, Output]

type ShiftLeft[Lhs, Rhs, Output any] Function2[Lhs, Rhs, Output]
type ShiftRight[Lhs, Rhs, Output any] Function2[Lhs, Rhs, Output]
