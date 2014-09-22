require 'sinatra'

post '*' do
	request.body.each_line.each_with_index do |line, i|
		print "#{i}: #{line}"
	end
end
