import { Item, ItemWithoutId } from "@/models/item";
import useFetch from "@/lib/useFetch";
import { auth } from "@clerk/nextjs";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { config } from "./config";

const storeURL = config.serverURL + "store/";

export function GetStaffByStore(storeId: string) {
  // const queryClient = useQueryClient()
  const authFetch = useFetch();

  console.log(storeURL + storeId + "/staff");
  return useQuery({
    queryKey: ["staff", storeId],
    queryFn: () =>
      authFetch({ url: storeURL + storeId + "/staff", toJSON: true }),
  });
}

export function SendInvite(storeId: string) {
  const authFetch = useFetch();

  return useMutation({
    mutationFn: (email: string) =>
      authFetch({
        url: `${storeURL}${storeId}/staff/invite`,
        init: {
          method: "POST",
          body: JSON.stringify({ email: email }),
        },
        headers: {
          "Content-Type": "application/json",
        },
      }),
    onError: (err, email, context) => {
      console.log("Error while sending invite to", email, ":", err);
    },
    onSuccess: (email) => {
      console.log("Invite sent successfully to", email);
      // You can add any additional logic you need on successful invite here
    },
  });
}