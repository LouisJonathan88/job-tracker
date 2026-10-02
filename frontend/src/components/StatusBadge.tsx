const STATUS_STYLES: Record<string, string> = {
  wishlist: "bg-gray-100 text-gray-700",
  applied: "bg-blue-100 text-blue-700",
  interview: "bg-yellow-100 text-yellow-700",
  offer: "bg-purple-100 text-purple-700",
  accepted: "bg-green-100 text-green-700",
  rejected: "bg-red-100 text-red-700",
};

export const STATUS_LABELS: Record<string, string> = {
  wishlist: "Wishlist",
  applied: "Melamar",
  interview: "Interview",
  offer: "Offer",
  accepted: "Diterima",
  rejected: "Ditolak",
};

export function StatusBadge({ status }: { status: string }) {
  const style = STATUS_STYLES[status] ?? "bg-gray-100 text-gray-700";
  const label = STATUS_LABELS[status] ?? status;

  return (
    <span className={`rounded-full px-3 py-1 text-xs font-medium ${style}`}>
      {label}
    </span>
  );
}