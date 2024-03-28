<!-- Start SDK Example Usage [usage] -->
```typescript
import { SDK } from "openapi";

async function run() {
    const sdk = new SDK();

    const result = await sdk.createDrink({
        ingredients: ["Gin", "Tonic water", "Lime juice"],
    });

    // Handle the result
    console.log(result);
}

run();

```
<!-- End SDK Example Usage [usage] -->