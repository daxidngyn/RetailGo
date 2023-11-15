"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm, SubmitHandler } from "react-hook-form";
import { useFetch } from "../../lib/utils"

import * as z from "zod";

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
import { Input } from "@/components/ui/input";

const formSchema = z.object({
  name: z.string(),
  price: z.coerce.number(),
  quantity: z.coerce.number(),
  category: z.string()
});

export default function AddItemDialog() {

  const form = useForm<z.infer <typeof formSchema>>({
    resolver: zodResolver(formSchema),
  });;

  let authFetch = useFetch()

  const onNewItem: SubmitHandler<z.infer<typeof formSchema>> = async (values: z.infer<typeof formSchema>) => {

    // console.log("Submit Triggered:", values)

    try {
      const response = await authFetch("http://localhost:8080/store/1391/inventory/create", 
        {
          method: 'POST',
          body: JSON.stringify(values, (key, value) => key === "quantity" || key === "price" ? parseFloat(value) : value)
        },
        {
          'Content-Type': 'application/json'
        }
      )
  
      if (!response.id) {
        throw new Error('Failed to save the item.');
      }
  
    } catch (error) {
      console.error("There was an error:", error);
    }
  };
  

  return (
    <Dialog>
      <DialogTrigger asChild>
        <button className="bg-amber-500 text-sm px-3 py-1.5 text-white font-medium rounded-md">
          Add item
        </button>
      </DialogTrigger>

      <Form {...form}>
        <form>
          <DialogContent>
            <form onSubmit = {form.handleSubmit(onNewItem, (data) => console.log("Error:", data))}>
            <DialogHeader>
              <DialogTitle>
                Add Item
              </DialogTitle>
            </DialogHeader>

            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Name</FormLabel>
                  <FormControl>
                    <Input {...field}/>
                  </FormControl>
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="category"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Category</FormLabel>
                  <FormControl>
                    <Input {...field} />
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
                    <Input {...field} type = "number"/>
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
                    <Input {...field} type = "number"/>
                  </FormControl>
                </FormItem>
              )}
            />

            <DialogFooter className="gap-x-4">
              <button type="submit">Save</button>
            </DialogFooter>
            </form>
          </DialogContent>
        </form>
      </Form>
    </Dialog>
  );
}
