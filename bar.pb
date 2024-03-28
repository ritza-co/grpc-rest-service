
3.0.0Q
Hitchhiker's Guide to Drinks*An API for creating and retrieving drinks.21.0.01
http://localhost:8080Local development server"�
�
/create-drink�2�0Create a new drink based on provided ingredients*createDrink:�
��

application/jsonk
+)
'#/components/schemas/IngredientsRequest<:ingredients:
    - Gin
    - Tonic water
    - Lime juice
B�
L
J
unexpected error6
4
application/json 

#/components/schemas/error�
200�
�
Successful response�
�
application/json�
&$
"#/components/schemas/DrinkResponse��name: Pan Galactic Gin and Tonic
description: A refreshing twist on the classic gin and tonic, inspired by the Hitchhiker's Guide to the Galaxy.
recipe: Fill a glass with ice. Pour in 2 oz of gin, 4 oz of tonic water, and a squeeze of lime juice. Stir gently and garnish with a lime wedge. Enjoy responsibly while contemplating the vastness of the universe.
photo: https://example.com/pan-galactic-gin-and-tonic.jpg

�

/get-drink�2�(Get a drink recipe by providing its name*getDrink:m
kg
e
application/jsonQ
)'
%#/components/schemas/DrinkNameRequest$"name: Pan Galactic Gargle Blaster
B�
L
J
unexpected error6
4
application/json 

#/components/schemas/error�
200�
�
Successful response�
�
application/json�
,*
(#/components/schemas/DrinkRecipeResponse��ingredients:
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
*�
�
�
IngredientsRequest�
��ingredients�object��
�
ingredientsr
p:#!- Gin
- Tonic water
- Lime juice
�array�

:Gin
�string�(An array of ingredient names as strings.�DA request containing a list of ingredients for creating a new drink.
�
DrinkResponse�
��object��
K
nameC
A:Pan Galactic Gin and Tonic
�string�The name of the drink.
�
description�
�:ecA refreshing twist on the classic gin and tonic, inspired by the Hitchhiker's Guide to the Galaxy.
�string�!A brief description of the drink.
�
recipe�
�:��Fill a glass with ice. Pour in 2 oz of gin, 4 oz of tonic water, and a squeeze of lime juice. Stir gently and garnish with a lime wedge. Enjoy responsibly while contemplating the vastness of the universe.
�string� The recipe for making the drink.
|
photos
q:53https://example.com/pan-galactic-gin-and-tonic.jpg
�string�(A URL pointing to an image of the drink.�uri�>A response containing information about a newly created drink.
�
DrinkNameRequest�
��name�object�N
L
nameD
B:Pan Galactic Gargle Blaster
�string�The name of the drink.�DA request containing the name of a drink to retrieve the recipe for.
�
DrinkRecipeResponse�
��object��
�
ingredients
}�array�-
+)
'#/components/schemas/IngredientQuantity�BAn array of objects representing ingredients and their quantities.
�
recipe�
�:��Pour the Ol' Janx Spirit into a glass, add the sea water and Acturan Mega-gin. Bubble the Fallian marsh gas through the mixture, then add the Qualactin Hypermint extract. Drink... but... very carefully.
�string� The recipe for making the drink.
}
photot
r:64https://example.com/pan-galactic-gargle-blaster.jpg
�string�(A URL pointing to an image of the drink.�uri�?A response containing the recipe and information about a drink.
�
IngredientQuantity�
��object��
E
name=
;:Ol' Janx Spirit
�string�The name of the ingredient.
O
quantityC
A:	1 bottle
�string�(The quantity of the ingredient required.�6An object representing an ingredient and its quantity.
^
errorU
S�code�message�object�6

code
�integer�int32

message
	�string