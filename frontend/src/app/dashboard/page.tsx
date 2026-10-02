"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import { useAuth } from "@/contexts/AuthContext";
import { getApplications } from "@/lib/api";
import type { Application } from "@/types";
import { StatusBadge, STATUS_LABELS } from "@/components/StatusBadge";
import { AppShell } from "@/components/AppShell";
import { Button } from "@/components/Button";

const STATUS_KEYS = Object.keys(STATUS_LABELS);

export default function DashboardPage() {
  const { isAuthenticated, isLoading } = useAuth();
  const router = useRouter();
  const [applications, setApplications] = useState<Application[]>([]);
  const [loadingData, setLoadingData] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [filter, setFilter] = useState<string | null>(null);

  useEffect(() => {
    if (!isLoading && !isAuthenticated) {
      router.push("/login");
    }
  }, [isLoading, isAuthenticated, router]);

  useEffect(() => {
    if (!isAuthenticated) return;

    getApplications()
      .then(setApplications)
      .catch((err) =>
        setError(err instanceof Error ? err.message : "Gagal memuat data"),
      )
      .finally(() => setLoadingData(false));
  }, [isAuthenticated]);

  if (isLoading || !isAuthenticated) {
    return null;
  }

  const counts: Record<string, number> = {};
  for (const status of STATUS_KEYS) {
    counts[status] = applications.filter(
      (app) => app.current_status === status,
    ).length;
  }

  const filteredApplications = filter
    ? applications.filter((app) => app.current_status === filter)
    : applications;

  function toggleFilter(status: string) {
    setFilter((current) => (current === status ? null : status));
  }

  return (
    <AppShell>
      <div className="mx-auto max-w-3xl">
        <div className="mb-6 flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
          <h1 className="text-2xl font-semibold text-gray-900">Dashboard</h1>
          <Link href="/applications/new">
            <Button className="w-full sm:w-auto">+ Tambah Lamaran</Button>
          </Link>
        </div>

        {/* Kartu statistik */}
        <div className="mb-8 grid grid-cols-2 gap-3 sm:grid-cols-3 md:grid-cols-4">
          <button
            onClick={() => setFilter(null)}
            className={`rounded-lg border p-4 text-left transition-colors ${
              filter === null
                ? "border-indigo-500 bg-indigo-50"
                : "bg-white hover:bg-gray-50"
            }`}
          >
            <p className="text-2xl font-semibold text-gray-900">
              {applications.length}
            </p>
            <p className="text-sm text-gray-600">Total Lamaran</p>
          </button>

          {STATUS_KEYS.map((status) => (
            <button
              key={status}
              onClick={() => toggleFilter(status)}
              className={`rounded-lg border p-4 text-left transition-colors ${
                filter === status
                  ? "border-indigo-500 bg-indigo-50"
                  : "bg-white hover:bg-gray-50"
              }`}
            >
              <p className="text-2xl font-semibold text-gray-900">
                {counts[status]}
              </p>
              <p className="text-sm text-gray-600">{STATUS_LABELS[status]}</p>
            </button>
          ))}
        </div>

        {/* Daftar lamaran */}
        {loadingData && <p className="text-gray-500">Memuat...</p>}
        {error && <p className="text-red-600">{error}</p>}

        {!loadingData && !error && filteredApplications.length === 0 && (
          <p className="text-gray-500">
            {filter
              ? `Tidak ada lamaran dengan status "${STATUS_LABELS[filter]}".`
              : "Belum ada lamaran."}
          </p>
        )}

        <div className="space-y-3">
          {filteredApplications.map((app) => (
            <Link
              key={app.id}
              href={`/applications/${app.id}`}
              className="block rounded-lg border bg-white p-4 shadow-sm hover:border-indigo-300"
            >
              <div className="flex items-center justify-between">
                <div>
                  <p className="font-medium text-gray-900">{app.position}</p>
                  <p className="text-sm text-gray-600">{app.company}</p>
                </div>
                <StatusBadge status={app.current_status} />
              </div>
            </Link>
          ))}
        </div>
      </div>
    </AppShell>
  );
}
