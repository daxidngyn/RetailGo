"use server";

import { revalidatePath } from "next/cache";

export const updateUser = async ({
  user,
  user_id,
}: {
  user: {
    first_name: string;
    last_name: string;
    email: string;
  };
  user_id: string;
}) => {
  const env: string = process.env.NODE_ENV;

  let serverUrl = `https://retailgo-production.up.railway.app/user/${user_id}/`;
  if (env === "development") {
    serverUrl = `http://localhost:8080/user/${user_id}/`;
  }
  console.log(`attempting to update user ${user_id}`);
  try {
    let response = await fetch(serverUrl, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(user),
    });
    console.log(`response: ${response.status} + ${response.statusText}`);
    if (response.status === 200) {
      revalidatePath("/user/[user]", "page");
      return true;
    }
  } catch (err) {
    console.error(`error updating user ${user_id}: ${err}`);
    return false;
  }
  return false;
};