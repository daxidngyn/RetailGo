import InventoryTable from "@/components/app-page/inventory-table";
import AddItemDialog from "@/components/app-page/item-dialog";
import { Item } from "@/models/item";
import { auth } from "@clerk/nextjs";
import { cx } from "class-variance-authority";

const stats = [
  {
    name: "Coming soon 1",
    value: "0",
    change: "+0%",
    changeType: "positive",
  },
  {
    name: "Coming soon 2",
    value: "0",
    change: "+0%",
    changeType: "positive",
  },
  {
    name: "Coming soon 3",
    value: "0",
    change: "+0%",
    changeType: "positive",
  },
];

export default async function Inventory({
  params,
}: {
  params: { store_id: string };
}) {
  const { sessionId } = auth();

  const res = await fetch(
    `https://retailgo-production.up.railway.app/store/${params.store_id}/inventory`,
    {
      cache: "force-cache",
      headers: {
        Authorization: `Bearer ${sessionId}`,
      },
    }
  );
  if (!res.ok) return <div>failed to fetch pos stuff</div>;

  const items = JSON.parse(await res.text());

  const newItem = new Item();
  const data = JSON.parse(JSON.stringify(newItem));

  return (
    <main className="bg-gray-50 h-full flex-grow">
      <div className="py-6 px-6 md:px-8 max-w-6xl mx-auto lg:ml-0">
        <div className="flex items-center justify-between ">
          <h1 className="text-2xl font-bold">Inventory</h1>

          <div>
            <AddItemDialog item={data} />
          </div>
        </div>
        <hr className="my-4" />

        <div className="mt-6">
          <dl className="border rounded-md mx-auto grid grid-cols-1 gap-px bg-gray-900/5 sm:grid-cols-2 lg:grid-cols-4">
            <div className="rounded-tl-md rounded-bl-md flex flex-wrap items-baseline justify-between gap-x-4 gap-y-2 bg-white px-4 py-10 sm:px-6 xl:px-8">
              <dt className="text-sm font-medium leading-6 text-gray-500">
                Total items
              </dt>
              {/* <dd className="text-gray-700 text-sm font-medium">
                  {stat.change}
                </dd> */}
              <dd className="w-full flex-none text-3xl font-medium leading-10 tracking-tight text-gray-900">
                {items.length}
              </dd>
            </div>

            {stats.map((stat, idx) => (
              <div
                key={stat.name}
                className={cx(
                  "flex flex-wrap items-baseline justify-between gap-x-4 gap-y-2 bg-white px-4 py-10 sm:px-6 xl:px-8",
                  idx === stats.length - 1 && "rounded-tr-md rounded-br-md"
                )}
              >
                <dt className="text-sm font-medium leading-6 text-gray-500">
                  {stat.name}
                </dt>
                <dd className="w-full flex-none text-3xl font-medium leading-10 tracking-tight text-gray-900">
                  {stat.value}
                </dd>
              </div>
            ))}
          </dl>
        </div>

        <div className="mt-6">
          <InventoryTable items={items} />
        </div>
      </div>
    </main>
  );
}