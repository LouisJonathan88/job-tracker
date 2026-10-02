import { ReactNode } from "react";

export function Card({ children, className = "" }: { children: ReactNode; className?: string }) {
  return (
    <div className={`rounded-lg border bg-white p-8 shadow-sm ${className}`}>
      {children}
    </div>
  );
}