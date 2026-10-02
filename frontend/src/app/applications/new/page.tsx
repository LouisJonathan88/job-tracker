// "use client";

// import { useState, FormEvent } from "react";
// import { useRouter } from "next/navigation";
// import { createApplication } from "@/lib/api";

// export default function NewApplicationPage() {
//   const router = useRouter();
//   const [company, setCompany] = useState("");
//   const [position, setPosition] = useState("");
//   const [jobUrl, setJobUrl] = useState("");
//   const [appliedAt, setAppliedAt] = useState("");
//   const [notes, setNotes] = useState("");
//   const [error, setError] = useState<string | null>(null);
//   const [isSubmitting, setIsSubmitting] = useState(false);

//   async function handleSubmit(e: FormEvent) {
//     e.preventDefault();
//     setError(null);
//     setIsSubmitting(true);

//     try {
//       await createApplication({
//         company,
//         position,
//         job_url: jobUrl || undefined,
//         applied_at: appliedAt || undefined,
//         notes: notes || undefined,
//       });
//       router.push("/dashboard");
//     } catch (err) {
//       setError(err instanceof Error ? err.message : "Gagal membuat lamaran");
//       setIsSubmitting(false);
//     }
//   }

//   return (
//     <div className="min-h-screen bg-gray-50 p-8">
//       <div className="mx-auto max-w-lg rounded-lg border bg-white p-8 shadow-sm">
//         <h1 className="mb-6 text-2xl font-semibold text-gray-900">Tambah Lamaran</h1>

//         <form onSubmit={handleSubmit} className="space-y-4">
//           <div>
//             <label htmlFor="company" className="block text-sm font-medium text-gray-700">
//               Perusahaan
//             </label>
//             <input
//               id="company"
//               type="text"
//               required
//               value={company}
//               onChange={(e) => setCompany(e.target.value)}
//               className="mt-1 w-full rounded-md border border-gray-300 px-3 py-2 text-gray-900 focus:border-blue-500 focus:outline-none"
//             />
//           </div>

//           <div>
//             <label htmlFor="position" className="block text-sm font-medium text-gray-700">
//               Posisi
//             </label>
//             <input
//               id="position"
//               type="text"
//               required
//               value={position}
//               onChange={(e) => setPosition(e.target.value)}
//               className="mt-1 w-full rounded-md border border-gray-300 px-3 py-2 text-gray-900 focus:border-blue-500 focus:outline-none"
//             />
//           </div>

//           <div>
//             <label htmlFor="jobUrl" className="block text-sm font-medium text-gray-700">
//               Link Lowongan (opsional)
//             </label>
//             <input
//               id="jobUrl"
//               type="url"
//               value={jobUrl}
//               onChange={(e) => setJobUrl(e.target.value)}
//               className="mt-1 w-full rounded-md border border-gray-300 px-3 py-2 text-gray-900 focus:border-blue-500 focus:outline-none"
//             />
//           </div>

//           <div>
//             <label htmlFor="appliedAt" className="block text-sm font-medium text-gray-700">
//               Tanggal Melamar (opsional)
//             </label>
//             <input
//               id="appliedAt"
//               type="date"
//               value={appliedAt}
//               onChange={(e) => setAppliedAt(e.target.value)}
//               className="mt-1 w-full rounded-md border border-gray-300 px-3 py-2 text-gray-900 focus:border-blue-500 focus:outline-none"
//             />
//           </div>

//           <div>
//             <label htmlFor="notes" className="block text-sm font-medium text-gray-700">
//               Catatan (opsional)
//             </label>
//             <textarea
//               id="notes"
//               rows={3}
//               value={notes}
//               onChange={(e) => setNotes(e.target.value)}
//               className="mt-1 w-full rounded-md border border-gray-300 px-3 py-2 text-gray-900 focus:border-blue-500 focus:outline-none"
//             />
//           </div>

//           {error && <p className="text-sm text-red-600">{error}</p>}

//           <div className="flex gap-3">
//             <button
//               type="submit"
//               disabled={isSubmitting}
//               className="flex-1 rounded-md bg-blue-600 py-2 text-white hover:bg-blue-700 disabled:opacity-50"
//             >
//               {isSubmitting ? "Menyimpan..." : "Simpan"}
//             </button>
//             <button
//               type="button"
//               onClick={() => router.push("/dashboard")}
//               className="flex-1 rounded-md border border-gray-300 py-2 text-gray-700 hover:bg-gray-100"
//             >
//               Batal
//             </button>
//           </div>
//         </form>
//       </div>
//     </div>
//   );
// }


"use client";

import { useState, FormEvent } from "react";
import { useRouter } from "next/navigation";
import { createApplication } from "@/lib/api";
import { AppShell } from "@/components/AppShell";
import { Card } from "@/components/Card";
import { Input } from "@/components/Input";
import { Button } from "@/components/Button";

export default function NewApplicationPage() {
  const router = useRouter();
  const [company, setCompany] = useState("");
  const [position, setPosition] = useState("");
  const [jobUrl, setJobUrl] = useState("");
  const [appliedAt, setAppliedAt] = useState("");
  const [notes, setNotes] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [isSubmitting, setIsSubmitting] = useState(false);

  async function handleSubmit(e: FormEvent) {
    e.preventDefault();
    setError(null);
    setIsSubmitting(true);

    try {
      await createApplication({
        company,
        position,
        job_url: jobUrl || undefined,
        applied_at: appliedAt || undefined,
        notes: notes || undefined,
      });
      router.push("/dashboard");
    } catch (err) {
      setError(err instanceof Error ? err.message : "Gagal membuat lamaran");
      setIsSubmitting(false);
    }
  }

  return (
    <AppShell>
      <div className="mx-auto max-w-lg">
        <Card>
          <h1 className="mb-6 text-2xl font-semibold text-gray-900">Tambah Lamaran</h1>

          <form onSubmit={handleSubmit} className="space-y-4">
            <Input
              id="company"
              label="Perusahaan"
              type="text"
              required
              value={company}
              onChange={(e) => setCompany(e.target.value)}
            />

            <Input
              id="position"
              label="Posisi"
              type="text"
              required
              value={position}
              onChange={(e) => setPosition(e.target.value)}
            />

            <Input
              id="jobUrl"
              label="Link Lowongan (opsional)"
              type="url"
              value={jobUrl}
              onChange={(e) => setJobUrl(e.target.value)}
            />

            <Input
              id="appliedAt"
              label="Tanggal Melamar (opsional)"
              type="date"
              value={appliedAt}
              onChange={(e) => setAppliedAt(e.target.value)}
            />

            <div>
              <label htmlFor="notes" className="block text-sm font-medium text-gray-700">
                Catatan (opsional)
              </label>
              <textarea
                id="notes"
                rows={3}
                value={notes}
                onChange={(e) => setNotes(e.target.value)}
                className="mt-1 w-full rounded-md border border-gray-300 px-3 py-2 text-gray-900 focus:border-indigo-500 focus:outline-none"
              />
            </div>

            {error && <p className="text-sm text-red-600">{error}</p>}

            <div className="flex gap-3">
              <Button type="submit" disabled={isSubmitting} className="flex-1">
                {isSubmitting ? "Menyimpan..." : "Simpan"}
              </Button>
              <Button
                type="button"
                variant="secondary"
                onClick={() => router.push("/dashboard")}
                className="flex-1"
              >
                Batal
              </Button>
            </div>
          </form>
        </Card>
      </div>
    </AppShell>
  );
}