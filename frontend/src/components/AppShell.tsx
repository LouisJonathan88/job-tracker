"use client";

import { ReactNode } from "react";
import { LogOut } from "lucide-react";
import { useAuth } from "@/contexts/AuthContext";

export function AppShell({ children }: { children: ReactNode }) {
  const { logout } = useAuth();

  return (
    <div className="min-h-screen bg-gray-50">
      <header className="border-b bg-white">
        <div className="mx-auto flex max-w-5xl items-center justify-between px-4 py-3 sm:px-6">
          <h2 className="text-lg font-semibold text-indigo-600">Job Tracker</h2>

          <nav className="flex items-center gap-1">
            <button
              onClick={logout}
              className="flex items-center gap-2 rounded-md px-3 py-2 text-sm font-medium text-gray-700 hover:bg-gray-100"
            >
              <LogOut size={18} />
              <span className="hidden sm:inline">Keluar</span>
            </button>
          </nav>
        </div>
      </header>

      <main className="mx-auto max-w-5xl p-4 sm:p-8">{children}</main>
    </div>
  );
}