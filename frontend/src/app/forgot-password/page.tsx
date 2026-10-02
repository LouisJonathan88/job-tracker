"use client";

import { useState, FormEvent } from "react";
import Link from "next/link";
import { forgotPassword } from "@/lib/api";
import { Card } from "@/components/Card";
import { Input } from "@/components/Input";
import { Button } from "@/components/Button";

export default function ForgotPasswordPage() {
  const [email, setEmail] = useState("");
  const [message, setMessage] = useState<string | null>(null);
  const [isSubmitting, setIsSubmitting] = useState(false);

  async function handleSubmit(e: FormEvent) {
    e.preventDefault();
    setIsSubmitting(true);

    try {
      const res = await forgotPassword(email);
      setMessage(res.message);
    } catch {
      setMessage("Kalau email terdaftar, kami sudah mengirim link reset password");
    } finally {
      setIsSubmitting(false);
    }
  }

  return (
    <div className="flex min-h-screen items-center justify-center bg-gray-50">
      <Card className="w-full max-w-sm">
        <h1 className="mb-2 text-2xl font-semibold text-gray-900">Lupa Password</h1>
        {!message && (
          <p className="mb-6 text-sm text-gray-600">
            Masukkan email akunmu, kami akan mengirim link untuk reset password.
          </p>
        )}

        {message ? (
          <p className="text-sm text-gray-700">{message}</p>
        ) : (
          <form onSubmit={handleSubmit} className="space-y-4">
            <Input
              id="email"
              label="Email"
              type="email"
              required
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />

            <Button type="submit" disabled={isSubmitting} className="w-full">
              {isSubmitting ? "Mengirim..." : "Kirim Link Reset"}
            </Button>
          </form>
        )}

        <Link href="/login" className="mt-4 inline-block text-sm text-indigo-600 hover:underline">
          Kembali ke login
        </Link>
      </Card>
    </div>
  );
}