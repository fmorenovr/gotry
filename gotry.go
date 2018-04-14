package gotry

// Gotry object
type Gotry struct {
  catch   func(Exception)
  finally func()
  Error   Exception
}

// Throw function (return or rethrow an exception)
func Throw(e Exception) {
  if e == nil {
    panic(gotry_rethrow)
  } else {
    panic(e)
  }
}

// try this function
func Try(funcToTry func()) (o *Gotry) {
  o = &Gotry{nil, nil, nil}
  // catch throw in try
  defer func() {
    o.Error = recover()
  }()
  // do the func
  funcToTry()
  return
}

// catch function
func (o *Gotry) Catch(funcCatched func(err Exception)) (*Gotry){
  o.catch = funcCatched
  if o.Error != nil {
    defer func() {
      // call finally
      if o.finally != nil {
        o.finally()
      }
      // rethrow Gotrys from catch
      if err := recover(); err != nil {
        if err == gotry_rethrow {
          err = o.Error
        }
        panic(err)
      }
    }()
    o.catch(o.Error)
  } else if o.finally != nil {
    o.finally()
  }
  return o
}

// finally function
func (o *Gotry) Finally(finallyFunc func()) () {
  if o.finally != nil {
    panic("Finally Function by default !!")
  } else {
    o.finally = finallyFunc
  }
  defer o.finally()
  return
}
