import { Item, ItemWithoutId } from "@/models/item";
import useFetch from "@/lib/useFetch";
import { auth } from "@clerk/nextjs";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { config } from "./config";

const storeURL = config.serverURL + "store/";

export function useCreateItem(store: string) {
  const queryClient = useQueryClient();
  const authFetch = useFetch();

  return useMutation({
    mutationFn: (newItem: ItemWithoutId) =>
      authFetch({
        url: config.serverURL + "create/store",
        init: {
          method: "POST",
          body: JSON.stringify(newItem, (key, value) =>
            key === "quantity" || key === "price" ? parseFloat(value) : value
          ),
        },
        headers: {
          "Content-Type": "application/json",
        },
      }),
    onMutate: (newItem: ItemWithoutId) => {
      // await queryClient.cancelQueries({ queryKey: ['todos'] })
      const prevItems = queryClient.getQueryData(["items", store]) as Item[];
      queryClient.setQueryData(["items", store], (old: Item[]) => {
        return prevItems.length > 0
          ? [
              ...prevItems,
              { ...newItem, id: prevItems.length + prevItems[0].id },
            ]
          : [{ ...newItem, id: 0 }];
      });
      return { prevItems };
    },
    onError: (err, newItem: ItemWithoutId, context) => {
      console.log("Error while creating", newItem.name, ":", err);
      queryClient.setQueryData(["items", store], context?.prevItems);
    },
    onSuccess: (newItem: ItemWithoutId) => {
      queryClient.invalidateQueries({ queryKey: ["items", store] });
    },
  });
}