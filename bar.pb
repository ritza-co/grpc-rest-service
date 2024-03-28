
3.0.0Q
Hitchhiker's Guide to Drinks*An API for creating and retrieving drinks.21.0.01
http://localhost:8080Local development server"ı
½
/create-drink«2¨0Create a new drink based on provided ingredients*createDrink:‰
†

application/jsonk
+)
'#/components/schemas/IngredientsRequest<:ingredients:
    - Gin
    - Tonic water
    - Lime juice
BÚ
L
J
unexpected error6
4
application/json 

#/components/schemas/error‰
200
ş
Successful responseæ
ã
application/jsonÎ
&$
"#/components/schemas/DrinkResponse£ name: Pan Galactic Gin and Tonic
description: A refreshing twist on the classic gin and tonic, inspired by the Hitchhiker's Guide to the Galaxy.
recipe: Fill a glass with ice. Pour in 2 oz of gin, 4 oz of tonic water, and a squeeze of lime juice. Stir gently and garnish with a lime wedge. Enjoy responsibly while contemplating the vastness of the universe.
photo: https://example.com/pan-galactic-gin-and-tonic.jpg

º

/get-drink«2¨(Get a drink recipe by providing its name*getDrink:m
kg
e
application/jsonQ
)'
%#/components/schemas/DrinkNameRequest$"name: Pan Galactic Gargle Blaster
B‚
L
J
unexpected error6
4
application/json 

#/components/schemas/error±
200©
¦
Successful response
‹
application/jsonö
,*
(#/components/schemas/DrinkRecipeResponseÅÂingredients:
    - name: Ol' Janx Spirit
      quantity: 1 bottle
    - name: Sea water from Santraginus V
      quantity: 1 measure
    - name: Acturan Mega-gin
      quantity: 3 cubes
    - name: Fallian marsh gas
      quantity: 4 liters
    - name: Qualactin Hypermint extract
      quantity: 1 teaspoon
recipe: Pour the Ol' Janx Spirit into a glass, add the sea water and Acturan Mega-gin. Bubble the Fallian marsh gas through the mixture, then add the Qualactin Hypermint extract. Drink... but... very carefully.
photo: https://example.com/pan-galactic-gargle-blaster.jpg
*š
—
€
IngredientsRequesté
æºingredientsÊobjectú„

ingredientsr
p:#!- Gin
- Tonic water
- Lime juice
Êarrayò

:Gin
Êstring’(An array of ingredient names as strings.’DA request containing a list of ingredients for creating a new drink.
è
DrinkResponseÖ
ÓÊobjectú…
K
nameC
A:Pan Galactic Gin and Tonic
Êstring’The name of the drink.
§
description—
”:ecA refreshing twist on the classic gin and tonic, inspired by the Hitchhiker's Guide to the Galaxy.
Êstring’!A brief description of the drink.

recipe‚
ÿ:ĞÍFill a glass with ice. Pour in 2 oz of gin, 4 oz of tonic water, and a squeeze of lime juice. Stir gently and garnish with a lime wedge. Enjoy responsibly while contemplating the vastness of the universe.
Êstring’ The recipe for making the drink.
|
photos
q:53https://example.com/pan-galactic-gin-and-tonic.jpg
Êstring’(A URL pointing to an image of the drink.šuri’>A response containing information about a newly created drink.
À
DrinkNameRequest«
¨ºnameÊobjectúN
L
nameD
B:Pan Galactic Gargle Blaster
Êstring’The name of the drink.’DA request containing the name of a drink to retrieve the recipe for.
ˆ
DrinkRecipeResponseğ
íÊobjectú

ingredients
}Êarrayò-
+)
'#/components/schemas/IngredientQuantity’BAn array of objects representing ingredients and their quantities.
‹
recipe€
ı:ÎËPour the Ol' Janx Spirit into a glass, add the sea water and Acturan Mega-gin. Bubble the Fallian marsh gas through the mixture, then add the Qualactin Hypermint extract. Drink... but... very carefully.
Êstring’ The recipe for making the drink.
}
photot
r:64https://example.com/pan-galactic-gargle-blaster.jpg
Êstring’(A URL pointing to an image of the drink.šuri’?A response containing the recipe and information about a drink.
ø
IngredientQuantityá
ŞÊobjectú˜
E
name=
;:Ol' Janx Spirit
Êstring’The name of the ingredient.
O
quantityC
A:	1 bottle
Êstring’(The quantity of the ingredient required.’6An object representing an ingredient and its quantity.
^
errorU
SºcodeºmessageÊobjectú6

code
Êintegeršint32

message
	Êstring