# Complains about Go

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

	
