
export type Plan = {
    PlanID: string;
    Name: string;
    Description: string;
    PeriodStart: string;
    PeriodEnd: string;
    Duration: string;
};
export type PlanItem = {
  PlanItemID: string;
  BookID: string;
  Book: {
    ID: string;
    Name: string;
    Osis: string;
    ShortName: string;
    Number: number;
    BookDivision: string;
    Testament: string;
    ChapterCount: number;
    Summary: string;
    CreatedAt: string;
    UpdatedAt: string;
    Slug: string;
    YearWritten: string;
    PlaceWritten: string;
    VerseCount: number;
    Writers: string;
    PeopleCount: number;
    PlaceCount: number;
  };
  ChapterID: string;
  Chapter: {
    ID: string;
    BookID: string;
    Osis: string;
    Number: number;
    Summary: string;
    Writer: string;
    Slug: string;
    PeopleCount: number;
    PlaceCount: number;
    WriterCount: number;
    CreatedAt: Date;
    UpdatedAt: Date;
  };
  ReadDate: string;
  Summary?: string;
  CompletedBy?: string[];
  CreatedAt: Date;
  UpdatedAt: Date;
};