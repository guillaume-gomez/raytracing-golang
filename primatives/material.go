package primatives

type Material interface {
  Bounce(input Ray, hit HitRecord) (bool, Ray)
  Color() Vector
}