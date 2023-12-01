"use client"
import { PostJoinStore } from "@/lib/hooks/user";
import { SignUp, SignedIn, SignedOut, useUser } from "@clerk/nextjs";
import { on } from "events";
import Link from "next/link";
import { useSearchParams } from 'next/navigation'
import { currentUser } from '@clerk/nextjs';


const searchParams = useSearchParams()
const store_uuid = searchParams.get('code') || "";
const postJoinStore = PostJoinStore(store_uuid);

export async function OnAccept() {
  const { user } = useUser();
  if (!user) return null;
  if(!user)
    postJoinStore.mutate(user.id);
}

export function OnDecline() {
  console.log("clicked")
}


const InvitePage = () => {
  return (
    <div className=" relative isolate px-6 pt-14 lg:px-8 flex-1 flex flex-col justify-center items-center">
      <SignedIn>
        <h1 className="text-center text-2xl">
          <span font-bold>"Stores are us"</span> has invited you to their business
        </h1>
    
        <div className="py-10 flex items-center justify-center space-x-5">
            <Link
              href="#"
              onClick={OnAccept}
              className="bg-green-500 hover:bg-green-600 active:bg-green-700 focus:outline-none focus:ring focus:ring-green-300 text-white font-medium px-3 py-3 rounded-md text-sm">
              Accept Invitation
            </Link>
            <Link
              href="#"
              onClick={OnDecline}
              className="bg-red-500 hover:bg-red-600 active:bg-red-700 focus:outline-none focus:ring focus:ring-red-300 text-white font-medium px-3 py-3 rounded-md text-sm">
              Decline Invitation
            </Link>
          </div>
      </SignedIn>
      <SignedOut>
        <div>You are signed Out</div>
      </SignedOut>

    </div>
  );
};

export default InvitePage;
