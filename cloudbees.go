message PurchaseRequest {
  string from = 1;
  string to = 2;
  string first_name = 3;
  string last_name = 4;
  string email = 5;
}


message PurchaseResponse {
  string receipt_id = 1;
  string from = 2;
  string to = 3;
  string user = 4;
  double price = 5;
  string seat = 6;
}

message ReceiptRequest {
  string receipt_id = 1;
}


message ReceiptResponse {
  string from = 1;
  string to = 2;
  string user = 3;
  double price = 4;
  string seat = 5;
}


message SectionRequest {
  string section = 1;
}


message SectionResponse {
  repeated string users = 1;
}


message RemoveRequest {
  string receipt_id = 1;
}


message RemoveResponse {
  bool success = 1;
}


message ModifyRequest {
  string receipt_id = 1;
  string seat = 2;
}


message ModifyResponse {
  bool success = 1;
}


service TrainTicket {
  
  rpc PurchaseTicket (PurchaseRequest) returns (PurchaseResponse) {}

  
  rpc GetReceipt (ReceiptRequest) returns (ReceiptResponse) {}

  
  rpc GetSectionUsers (SectionRequest) returns (SectionResponse) {}

  
  rpc RemoveUser (RemoveRequest) returns (RemoveResponse) {}

  
  rpc ModifySeat (ModifyRequest) returns (ModifyResponse) {}
}


import (
  "context"
  "fmt"
  "log"
  "net"

  "google.golang.org/grpc"
  pb "path/to/your/protobuf/file"
)


type server struct {
  pb.UnimplementedTrainTicketServer
  
  receipts map[string]*pb.PurchaseResponse
  
  seats map[string][]string
}


func (s *server) PurchaseTicket(ctx context.Context, in *pb.PurchaseRequest) (*pb.PurchaseResponse, error) {
  
  receipt_id := fmt.Sprintf("%d", rand.Intn(1000000))
  
  price := 20.0
  
  seat := s.allocateSeat()
  
  response := &pb.PurchaseResponse{
    ReceiptId: receipt_id,
    From: in.From,
    To: in.To,
    User: fmt.Sprintf("%s %s", in.FirstName, in.LastName),
    Price: price,
    Seat: seat,
  }
  
  s.receipts[receipt_id] = response
  
  return response, nil
}


func (s *server) GetReceipt(ctx context.Context, in *pb.ReceiptRequest) (*pb.ReceiptResponse, error) {
  
  receipt, ok := s.receipts[in.ReceiptId]
  if !ok {
    
    return nil, fmt.Errorf("receipt id %s not found", in.ReceiptId)
  }
  
  response := &pb.ReceiptResponse{
    From: receipt.From,
    To: receipt.To,
    User: receipt.User,
    Price: receipt.Price,
    Seat: receipt.Seat,
  }
  
  return response, nil
}


func (s *server) GetSectionUsers(ctx context.Context, in *pb.SectionRequest) (*pb.SectionResponse, error) {
  
  seats, ok := s.seats[in.Section]
  if !ok {
    
    return nil, fmt.Errorf("section %s not valid", in.Section)
  }
  
  users := []string{}
  
  for _, seat := range seats {
    for _, receipt := range s.receipts {
      if receipt.Seat == seat {
        users = append(users, receipt.User)
        break
      }
    }
  }
  
  response := &pb.SectionResponse{
    Users: users,
  }
  
  return response, nil
}


func (s *server) RemoveUser(ctx context.Context, in *pb.RemoveRequest) (*pb.RemoveResponse, error) {
  
  receipt, ok := s.receipts[in.ReceiptId]
  if !ok {
    
    return nil, fmt.Errorf("receipt id %s not found", in.ReceiptId)
  }
  
  seat := receipt.Seat
  
  section := seat[:1]
  
  s.seats[section] = remove(s.seats[section], seat)
  
  delete(s.receipts, in.ReceiptId)
  
  response := &pb.RemoveResponse{
    Success: true,
  }
  
  return response, nil
}


func (s *server) ModifySeat(ctx context.Context, in *pb.ModifyRequest) (*pb.ModifyResponse, error) {
  
  receipt, ok := s.receipts[in.ReceiptId]
  if !ok {
    
    return nil, fmt.Errorf("receipt id %s not found", in.ReceiptId)
  }
  
  new_seat := in.Seat
  
  if !s.isValidSeat(new_seat) || !s.isAvailableSeat(new_seat) {
    
    return nil, fmt.Errorf("seat %s not valid or available", new_seat)
  }
  
  old_seat := receipt.Seat
  
  old_section := old_seat[:1]
  
  new_section := new_seat[:1]
  
  s.seats[old_section] = remove(s.seats[old_section], old_seat)
  
  s.seats[new_section] = append(s.seats[new_section], new_seat)
  
  receipt.Seat = new_seat
  
  response := &pb.ModifyResponse{
    Success: true,
  }
  
  return response, nil
}


func (s *server) allocateSeat() string {
  
  sections := []string{"A", "B"}
  seats_per_section := 10
  
  for _, section := range sections {
    
    for i := 1; i <= seats_per_section; i++ {
      
      seat := fmt.Sprintf("%s%d", section, i)
      
      if s.isAvailableSeat(seat) {
        
        return seat
      }
    }
  }
  
  return ""
}


func (s *server) isValidSeat(seat string) bool {
  // Define the valid sections and the seats per section
  valid_sections := []string{"A", "B"}
  seats_per_section := 10
  
  if seat == "" {
    
    return false
  }
  
  section := seat[:1]
  number,
