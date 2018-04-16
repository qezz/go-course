# Complains about Go

> The most comprehensive thoughs are already told by Sylvain Wallez,
> see their article "[Go: the Good, the Bad and the Ugly](https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/)"

1. **Zero-values are bullshit!** There is a bunch of cases when it isn't
   possible to think through the default value. For example, if you
   create a Point, it **mustn't** be placed at (0, 0)!
   * It leads to problems when you want to add some kind of new field
     to a struct, but you can't determine the place you have forgot to
     initialize it!
2. **Two kinds of variable declaration are weird (bullshit)!** As the
   result of zero-values, there is a declaration without
   initialization that syntatically differs from short
   declaration. It's an effing nightmare.
3. **Dependency manager is a skrewdriver for nails!** And you already
   know it.
4. **Type conversions are not explicit enough!** For compiler it's OK
   to have smth like `2 * 0.5`, but 

    ```go
    type A struct {
        field int64
    }
    A.field * 0.5
    ```

   yields a compile time error!
5. **No algebraic types!** Imagine you want to load a value from a
   map. You should do it like this:

    ```go
	val, ok := m["key"]
	if ok {
		// do smth with val
	}
	
	// or 
	if !ok {
		// return;
	}	
	```
   
   This error checking makes your code to look like spaghetti. For
   example, in rust you receive an `Option<T>` type, therefore `val`
   to be `Some(value)` or `None` and you can't forget to check it for
   existance.
6. **No practical Error type!** This issue makes error handling look like
	smth you can't control: it is whether *it* (`string`) or `nil`.
	
