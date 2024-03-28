/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import * as z from "zod";

/**
 * An object representing an ingredient and its quantity.
 */
export type IngredientQuantity = {
    /**
     * The name of the ingredient.
     */
    name?: string | undefined;
    /**
     * The quantity of the ingredient required.
     */
    quantity?: string | undefined;
};

/** @internal */
export namespace IngredientQuantity$ {
    export type Inbound = {
        name?: string | undefined;
        quantity?: string | undefined;
    };

    export const inboundSchema: z.ZodType<IngredientQuantity, z.ZodTypeDef, Inbound> = z
        .object({
            name: z.string().optional(),
            quantity: z.string().optional(),
        })
        .transform((v) => {
            return {
                ...(v.name === undefined ? null : { name: v.name }),
                ...(v.quantity === undefined ? null : { quantity: v.quantity }),
            };
        });

    export type Outbound = {
        name?: string | undefined;
        quantity?: string | undefined;
    };

    export const outboundSchema: z.ZodType<Outbound, z.ZodTypeDef, IngredientQuantity> = z
        .object({
            name: z.string().optional(),
            quantity: z.string().optional(),
        })
        .transform((v) => {
            return {
                ...(v.name === undefined ? null : { name: v.name }),
                ...(v.quantity === undefined ? null : { quantity: v.quantity }),
            };
        });
}
