"use client";

import { useState } from "react";
import { useParams } from "next/navigation";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import * as z from "zod";
import {
  createItem,
  updateItem,
} from "@/app/(app-page)/store/[store_id]/inventory/actions";
import { cn, wait } from "@/lib/utils";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
} from "@/components/ui/form";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { ChevronsUpDown, PencilIcon } from "lucide-react";

const formSchema = z.object({
  name: z.string(),
  price: z.coerce.number(),
  quantity: z.coerce.number(),
  category_name: z.string(),
});

export default function ItemDialog({
  item,
  mode = "add",
  categories,
}: {
  item?: Item;
  mode?: string;
  categories: Category[];
}) {
  const params = useParams();
  const [open, setOpen] = useState(false);

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: item ? item.name : undefined,
      price: item ? item.price : undefined,
      quantity: item ? item.quantity : undefined,
      category_name: item ? item.category_name : undefined,
    },
  });

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    if (!params.store_id) return;

    if (mode === "edit") {
      if (!item) return;

      await updateItem({
        item: { id: item.id, ...values },
        store_id: params.store_id as string,
      });

      wait().then(() => setOpen(false));
      return;
    }

    await createItem({ item: values, store_id: params.store_id as string });
    wait().then(() => setOpen(false));
  };

  const displayCategory = (value: string) => {
    if (!value) return "Select category";
    if (!categories || !categories.length) return value;

    const categoryName = categories.find((category) => category.name === value)
      ?.name;

    if (categoryName) return categoryName;

    return value;
  };

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        {mode === "edit" ? (
          <button className="icon-button">
            <PencilIcon className="h-5 w-5 p-0 text-amber-500" />
          </button>
        ) : (
          <button className="rounded-md bg-sky-500 px-4 py-1.5 text-sm font-medium text-white shadow-sm dark:bg-gradient-to-r dark:from-blue-600 dark:to-indigo-600">
            Add item
          </button>
        )}
      </DialogTrigger>

      <DialogContent className="dark:border-zinc-900 dark:bg-zinc-950">
        <DialogHeader>
          <DialogTitle>{mode === "edit" ? "Edit" : "Add"} Item</DialogTitle>
        </DialogHeader>

        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-2">
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Name</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      placeholder="Item name"
                      className="dark:border-zinc-800 dark:focus:ring-zinc-700"
                    />
                  </FormControl>
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="price"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Price</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      type="number"
                      placeholder="Item price"
                      className="dark:border-zinc-800 dark:focus:ring-zinc-700"
                    />
                  </FormControl>
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="quantity"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Quantity</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      type="number"
                      placeholder="Item quantity"
                      className="dark:border-zinc-800 dark:focus:ring-zinc-700"
                    />
                  </FormControl>
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="category_name"
              render={({ field }) => (
                <FormItem className="group pt-2">
                  <Popover>
                    <PopoverTrigger asChild>
                      <FormControl>
                        <Button
                          variant="outline"
                          role="combobox"
                          className={cn(
                            "w-full justify-between dark:border-zinc-800 dark:hover:bg-zinc-800",
                            !field.value && "text-muted-foreground",
                          )}
                        >
                          {displayCategory(field.value)}
                          <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
                        </Button>
                      </FormControl>
                    </PopoverTrigger>

                    <PopoverContent className="w-72 p-2 dark:border-zinc-800 dark:bg-zinc-900">
                      <Input
                        {...field}
                        placeholder="Enter a category"
                        className="dark:border-zinc-800 dark:focus:ring-zinc-700"
                      />

                      {categories.length ? (
                        <div className="mt-4 flex flex-col space-y-0.5">
                          {categories.map((category) => (
                            <Button
                              key={category.id}
                              variant="default"
                              value={category.name}
                              // @ts-ignore
                              onClick={(e) => field.onChange(e.target.value)}
                              className="justify-start bg-white text-left text-black shadow-none hover:bg-gray-100 hover:text-black dark:bg-zinc-800 dark:text-zinc-50"
                            >
                              {category.name}
                            </Button>
                          ))}
                        </div>
                      ) : (
                        <div className="pt-2 text-center">
                          <span className="text-sm text-muted-foreground">
                            You don{"'"}t have any categories created yet.
                            Create a new one now!
                          </span>
                        </div>
                      )}
                    </PopoverContent>
                  </Popover>
                </FormItem>
              )}
            />

            <DialogFooter className="gap-x-4 pt-2">
              <Button type="submit">
                {mode === "edit" ? "Update" : "Add"}
              </Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
}
