"use client";

import { useEffect, useState } from "react";
import { useRouter, useParams } from "next/navigation";
import Link from "next/link";
import {
  getApplication,
  updateApplicationStatus,
  deleteApplication,
} from "@/lib/api";
import type { Application } from "@/types";
import { StatusBadge } from "@/components/StatusBadge";
import { AppShell } from "@/components/AppShell";
import { Card } from "@/components/Card";
import { Button } from "@/components/Button";
import { Trash2 } from "lucide-react";
import { ConfirmDialog } from "@/components/ConfirmDialog";

const STATUS_OPTIONS: Record<string, string[]> = {
  wishlist: ["applied", "rejected"],
  applied: ["interview", "rejected"],
  interview: ["offer", "rejected"],
  offer: ["accepted", "rejected"],
  accepted: [],
  rejected: [],
};

function formatDate(dateString: string): string {
  return new Date(dateString).toLocaleDateString("id-ID", {
    day: "numeric",
    month: "long",
    year: "numeric",
  });
}

export default function ApplicationDetailPage() {
  const router = useRouter();
  const params = useParams();
  const id = Number(params.id);

  const [application, setApplication] = useState<Application | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [isUpdating, setIsUpdating] = useState(false);

  const [showDeleteConfirm, setShowDeleteConfirm] = useState(false);

  useEffect(() => {
    getApplication(id)
      .then(setApplication)
      .catch((err) =>
        setError(err instanceof Error ? err.message : "Gagal memuat data"),
      )
      .finally(() => setLoading(false));
  }, [id]);

  async function handleStatusChange(newStatus: string) {
    setIsUpdating(true);
    setError(null);
    try {
      const updated = await updateApplicationStatus(id, { status: newStatus });
      setApplication(updated);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Gagal mengubah status");
    } finally {
      setIsUpdating(false);
    }
  }

  async function confirmDelete() {
    setShowDeleteConfirm(false);
    try {
      await deleteApplication(id);
      router.push("/dashboard");
    } catch (err) {
      setError(err instanceof Error ? err.message : "Gagal menghapus lamaran");
    }
  }

  if (loading) {
    return (
      <AppShell>
        <p className="text-gray-500">Memuat...</p>
      </AppShell>
    );
  }

  if (!application) {
    return (
      <AppShell>
        <p className="text-red-600">{error || "Lamaran tidak ditemukan"}</p>
      </AppShell>
    );
  }

  const availableStatuses = STATUS_OPTIONS[application.current_status] ?? [];

  return (
    <AppShell>
      <div className="mx-auto max-w-lg">
        <Link
          href="/dashboard"
          className="mb-4 inline-block text-sm text-indigo-600 hover:underline"
        >
          Kembali ke Dashboard
        </Link>

        <Card>
          <div className="mb-4 flex items-start justify-between">
            <div>
              <h1 className="text-xl font-semibold text-gray-900">
                {application.position}
              </h1>
              <p className="text-gray-600">{application.company}</p>
            </div>
            <StatusBadge status={application.current_status} />
          </div>

          {application.job_url && (
            <p className="mb-2 text-sm">
              <a
                href={application.job_url}
                target="_blank"
                rel="noopener noreferrer"
                className="text-indigo-600 hover:underline"
              >
                Lihat lowongan
              </a>
            </p>
          )}

          {application.applied_at && (
            <p className="mb-2 text-sm text-gray-600">
              Dilamar pada: {formatDate(application.applied_at)}
            </p>
          )}

          {application.notes && (
            <p className="mb-4 text-sm text-gray-700">{application.notes}</p>
          )}

          {error && <p className="mb-4 text-sm text-red-600">{error}</p>}

          {availableStatuses.length > 0 && (
            <div className="mb-4">
              <label className="block text-sm font-medium text-gray-700">
                Ubah status ke:
              </label>
              <div className="mt-2 flex flex-wrap gap-2">
                {availableStatuses.map((status) => (
                  <Button
                    key={status}
                    variant="secondary"
                    onClick={() => handleStatusChange(status)}
                    disabled={isUpdating}
                  >
                    {status}
                  </Button>
                ))}
              </div>
            </div>
          )}

          <button
            onClick={() => setShowDeleteConfirm(true)}
            className="rounded-md p-2 text-red-600 hover:bg-red-50"
            aria-label="Hapus lamaran"
          >
            <Trash2 size={18} />
          </button>
        </Card>
      </div>
      <ConfirmDialog
        isOpen={showDeleteConfirm}
        title="Hapus Lamaran"
        message="Yakin ingin menghapus lamaran ini?"
        onConfirm={confirmDelete}
        onCancel={() => setShowDeleteConfirm(false)}
      />
    </AppShell>
  );
}
