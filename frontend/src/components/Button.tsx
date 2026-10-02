import { ButtonHTMLAttributes } from "react";

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: "primary" | "secondary" | "danger";
}

export function Button({ variant = "primary", className = "", ...props }: ButtonProps) {
  const base = "rounded-md px-4 py-2 text-sm font-medium transition-colors disabled:opacity-50";

  const variants = {
    primary: "bg-indigo-600 text-white hover:bg-indigo-700",
    secondary: "border border-gray-300 text-gray-700 hover:bg-gray-100",
    danger: "text-red-600 hover:underline",
  };

  return (
    <button className={`${base} ${variants[variant]} ${className}`} {...props} />
  );
}