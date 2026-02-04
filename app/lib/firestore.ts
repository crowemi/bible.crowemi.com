import { Firestore } from "@google-cloud/firestore";

const db = new Firestore({
  projectId: "crowemi-io-dev",
  databaseId: "crowemi-io-dev",
});

export type Filter = {
  Path: string;
  Value: string;
  Operator: "==" | "<" | ">" | "<=" | ">=";
};

interface Document {
  id: string;
}

export async function get<T extends Document>(
  collection: string,
  filters?: Filter[]
): Promise<T[]> {
  let query: FirebaseFirestore.Query = db.collection(collection);
  if (filters) {
    filters.forEach((filter) => {
      query = query.where(filter.Path, filter.Operator, filter.Value);
    });
  }
  const snapshot = await query.get();
  const ret: T[] = snapshot.docs.map((doc) => {
    const r = doc.data() as T;
    r.id = doc.id;
    return r;
  });
  return ret;
}
