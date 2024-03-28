/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import * as z from "zod";

/**
 * A request containing the name of a drink to retrieve the recipe for.
 */
export type DrinkNameRequest = {
    /**
     * The name of the drink.
     */
    name: string;
};

/** @internal */
export namespace DrinkNameRequest$ {
    export type Inbound = {
        name: string;
    };

    export const inboundSchema: z.ZodType<DrinkNameRequest, z.ZodTypeDef, Inbound> = z
        .object({
            name: z.string(),
        })
        .transform((v) => {
            return {
                name: v.name,
            };
        });

    export type Outbound = {
        name: string;
    };

    export const outboundSchema: z.ZodType<Outbound, z.ZodTypeDef, DrinkNameRequest> = z
        .object({
            name: z.string(),
        })
        .transform((v) => {
            return {
                name: v.name,
            };
        });
}
