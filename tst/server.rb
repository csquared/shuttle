require 'sinatra'

post '*' do
  puts request.body.read
end
