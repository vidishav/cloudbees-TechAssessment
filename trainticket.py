import Flask,request,json

app = Flask(-name-)
tickets = []

app.route('/purchase_ticket', methods=['POST'])
def purchase_ticket():
   data = request.json
    user = data.get('user')
    section = allocate_seat()
    ticket = {
        from: data.get('from'),
        to: data.get('to'),
        user: user,
        price_paid: 20,
           
section: section
    }
    tickets.append(ticket)
    return jsonify(ticket)

app.route('/receipt/<username>', methods=['GET'])
def get_receipt(username):
    user_tickets = [ticket for ticket in tickets if ticket['user']['username'] == username]
    return jsonify(user_tickets)

app.route('/view_users_by_section/<section>', methods=['GET'])
def view_users_by_section(section):
    users_in_section = [{'user': ticket['user'], 'section': ticket['section']} for ticket in tickets if ticket['section'] == section]
    return jsonify(users_in_section)

app.route('/remove_user/<username>', methods=['DELETE'])
def remove_user(username):
    global tickets
    tickets = [ticket for ticket in tickets if ticket['user']['username'] != username]
    return jsonify({'message': 'User removed successfully'})

app.route('/modify_seat/<username>', methods=['PUT'])
def modify_seat(username):
    data = request.json
    for ticket in tickets:
        if ticket['user']['username'] == username:
            ticket['section'] = data.get('section')
            return jsonify({'message': 'Seat modified successfully'})
    return jsonify({'error': 'User not found'})

def allocate_seat():
    
    count_a = sum(1 for ticket in tickets if ticket['section'] == 'A')
    count_b = sum(1 for ticket in tickets if ticket['section'] == 'B')
    return 'A' if count_a <= count_b else 'B'

if __name__ == '__main__':
    app.run(debug=True)
