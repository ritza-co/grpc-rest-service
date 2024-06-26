/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import * as z from "zod";

/**
 * A request containing a list of ingredients for creating a new drink.
 */
export type IngredientsRequest = {
    /**
     * An array of ingredient names as strings.
     */
    ingredients: Array<string>;
};

/** @internal */
export namespace IngredientsRequest$ {
    export type Inbound = {
        ingredients: Array<string>;
    };

    export const inboundSchema: z.ZodType<IngredientsRequest, z.ZodTypeDef, Inbound> = z
        .object({
            ingredients: z.array(z.string()),
        })
        .transform((v) => {
            return {
                ingredients: v.ingredients,
            };
        });

    export type Outbound = {
        ingredients: Array<string>;
    };

    export const outboundSchema: z.ZodType<Outbound, z.ZodTypeDef, IngredientsRequest> = z
        .object({
            ingredients: z.array(z.string()),
        })
        .transform((v) => {
            return {
                ingredients: v.ingredients,
            };
        });
}
