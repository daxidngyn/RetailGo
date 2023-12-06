import { Button } from "@/components/ui/button";

const POSOrderSummary = ({
  cart,
  TAX_RATE,
}: {
  cart: (Item & { quantity: number })[];
  TAX_RATE: number;
}) => {
  const subtotal = cart.reduce(
    (acc, val) => acc + parseFloat((val.price * val.quantity).toFixed(2)),
    0
  );

  const total = parseFloat((subtotal * TAX_RATE).toFixed(2));

  return (
    <div className="h-full">
      {cart.length ? (
        <div className="space-y-2">
          {cart.map((cartItem, idx) => (
            <div
              key={cartItem.id}
              className="flex items-center justify-between bg-gray-100 rounded-lg p-3"
            >
              <div className="flex items-center gap-x-2 text-sm">
                <div className="rounded-full bg-black text-gray-50 w-5 h-5 flex items-center justify-center">
                  <span className="text-xs">{idx + 1}</span>
                </div>
                <span>{cartItem.name}</span>
                <span className="text-gray-600">x{cartItem.quantity}</span>
              </div>

              <div>
                <span className="text-sm">
                  ${(cartItem.price * cartItem.quantity).toFixed(2)}
                </span>
              </div>
            </div>
          ))}
        </div>
      ) : (
        <div className="bg-gray-100 rounded-lg p-3">
          No items added to cart.
        </div>
      )}

      <div className="mt-6 p-3 bg-gray-100 rounded-lg">
        <div>
          <div className="flex items-center justify-between">
            <span className="text-sm text-gray-700">Subtotal</span>
            <span className="text-sm">${subtotal.toFixed(2)}</span>
          </div>

          <div className="flex items-center justify-between mt-1">
            <span className="text-sm text-gray-700">Tax</span>
            <span className="text-sm">
              $
              {(
                subtotal - parseFloat((subtotal * TAX_RATE).toFixed(2))
              ).toFixed(2)}
            </span>
          </div>

          <hr className="my-3" />

          <div className="flex items-center justify-between mt-1">
            <span className="text-gray-700">Total</span>
            <span className="font-medium">${total.toFixed(2)}</span>
          </div>

          <div className="mt-6">
            <Button className="rounded-full w-full py-5">Place order</Button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default POSOrderSummary;